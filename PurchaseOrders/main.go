package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

func main() {

	prefix := "main"

	log.Printf("%s: program starting\n", prefix)

	// Read all the files in the "data" subdirectory
	files, err := os.ReadDir("data")
	if err != nil {
		fmt.Printf("%s: {ERROR} %v\n", prefix, err)
	}

	// Now iterate through each file
	for _, file := range files {

		// If the file isn't a directory, read it as a text file
		if !file.IsDir() {

			log.Printf("%s: {INFO} reading file %q\n", prefix, file.Name())

			inputFile, err := os.ReadFile(filepath.Join("./data", file.Name()))
			if err != nil {
				log.Fatalf("%s: {ERROR} could not load foreign input file %v\n", prefix, err)
			}

			// Uncomment the next line to see the JSON data from the input file
			//log.Printf("%s: {INFO} file:\n%s\n", prefix, string(inputFile))

			// Let's unpack the JSON into our purchase order structure
			var my850 EDI850
			err = json.Unmarshal(inputFile, &my850)
			if err != nil {
				log.Fatalf("%s: {ERROR} could not unmarshal input file %v\n", prefix, err)
			}

			// Here is a typical example of something we would do with a purchase order: Iterate through it's line items
			// We will keep a running total of the PO price, and add that data element to the Purchase Order Object

			var poTotal float64
			for _, poLine := range my850.TransactionSets[0].Detail.BaselineItemDataPO1Loop {

				// We need to fill in some prices here, because they are always zero in OSD (I will explain during our class time)
				poLine.BaselineItemDataPO1.UnitPrice04 = math.Round(100*(100*rand.Float64()+15)) / 100.0

				log.Printf("%s: Line: %s Part: %s  Price: %5.2f Quantity: %4.1f \n",
					prefix, poLine.BaselineItemDataPO1.AssignedIdentification01, poLine.BaselineItemDataPO1.ProductServiceID07,
					poLine.BaselineItemDataPO1.UnitPrice04, poLine.BaselineItemDataPO1.QuantityOrdered02)

				poTotal += poLine.BaselineItemDataPO1.UnitPrice04 * poLine.BaselineItemDataPO1.QuantityOrdered02
			}

			// Now we need to add the total to the Purchase Order.  This is done through the MonetaryAmountAMT structure
			// This creates one that is labeled as a cumulative total and a debit
			myAMP := MonetaryAmountAMT{AmountQualifierCode01: "AA2", MonetaryAmount02: poTotal, CreditDebitFlagCode03: "D"}

			// If we don't have an existing CTT loop, let's create one -- don't want to reference a nil pointer
			if len(my850.TransactionSets[0].Summary.TransactionTotalsCTTLoop) == 0 {
				log.Printf("%s: {INFO} adding a new CTTLoop\n", prefix)
				my850.TransactionSets[0].Summary.TransactionTotalsCTTLoop = append(my850.TransactionSets[0].Summary.TransactionTotalsCTTLoop, TransactionTotalsCTTLoop{})
			} else {
				log.Printf("%s: {INFO} using an existing CTTLoop", prefix)
			}

			// Now we can set the AMT field safely, because it has memory allocated to it either way
			my850.TransactionSets[0].Summary.TransactionTotalsCTTLoop[0].MonetaryAmountAMT = &myAMP
		}
	}

	fmt.Printf("%s: {INFO} program terminating\n", prefix)
}

package main

type EDI850 struct {
	Envelope        Envelope          `json:"envelope,omitempty"`
	TransactionSets []TransactionSets `json:"transactionSets,omitempty"`
	Delimiters      Delimiters        `json:"delimiters,omitempty"`
}

type GroupHeader struct {
	AgencyCode               string `json:"agencyCode,omitempty"`
	ApplicationReceiverCode  string `json:"applicationReceiverCode,omitempty"`
	ApplicationSenderCode    string `json:"applicationSenderCode,omitempty"`
	ControlNumber            string `json:"controlNumber,omitempty"`
	Date                     string `json:"date,omitempty"`
	FunctionalIdentifierCode string `json:"functionalIdentifierCode,omitempty"`
	Release                  string `json:"release,omitempty"`
	Time                     string `json:"time,omitempty"`
}
type GroupTrailer struct {
	ControlNumber        string `json:"controlNumber,omitempty"`
	NumberOfTransactions string `json:"numberOfTransactions,omitempty"`
}
type InterchangeHeader struct {
	AcknowledgementRequestedCode      string `json:"acknowledgementRequestedCode,omitempty"`
	AuthorizationInformation          string `json:"authorizationInformation,omitempty"`
	AuthorizationInformationQualifier string `json:"authorizationInformationQualifier,omitempty"`
	ComponentSeparator                string `json:"componentSeparator,omitempty"`
	ControlNumber                     string `json:"controlNumber,omitempty"`
	ControlVersionNumber              string `json:"controlVersionNumber,omitempty"`
	Date                              string `json:"date,omitempty"`
	ReceiverID                        string `json:"receiverId,omitempty"`
	ReceiverQualifier                 string `json:"receiverQualifier,omitempty"`
	RepetitionSeparator               string `json:"repetitionSeparator,omitempty"`
	SecurityInformation               string `json:"securityInformation,omitempty"`
	SecurityQualifier                 string `json:"securityQualifier,omitempty"`
	SenderID                          string `json:"senderId,omitempty"`
	SenderQualifier                   string `json:"senderQualifier,omitempty"`
	Time                              string `json:"time,omitempty"`
	UsageIndicatorCode                string `json:"usageIndicatorCode,omitempty"`
}
type InterchangeTrailer struct {
	ControlNumber            string `json:"controlNumber,omitempty"`
	NumberOfFunctionalGroups string `json:"numberOfFunctionalGroups,omitempty"`
}

type Envelope struct {
	GroupHeader        GroupHeader        `json:"groupHeader,omitempty"`
	GroupTrailer       GroupTrailer       `json:"groupTrailer,omitempty"`
	InterchangeHeader  InterchangeHeader  `json:"interchangeHeader,omitempty"`
	InterchangeTrailer InterchangeTrailer `json:"interchangeTrailer,omitempty"`
}

type TransactionSets struct {
	Heading Heading `json:"heading,omitempty"`
	Detail  Detail  `json:"detail,omitempty"`
	Summary Summary `json:"summary,omitempty"`
}
type Delimiters struct {
	Element   string `json:"element,omitempty"`
	Composite string `json:"composite,omitempty"`
	Segment   string `json:"segment,omitempty"`
}

type TransactionSetHeaderST struct {
	TransactionSetIdentifierCode01 string `json:"transaction_set_identifier_code_01,omitempty"`
	TransactionSetControlNumber02  int    `json:"transaction_set_control_number_02,omitempty"`
}
type BeginningSegmentForPurchaseOrderBEG struct {
	TransactionSetPurposeCode01 string `json:"transaction_set_purpose_code_01,omitempty"`
	PurchaseOrderTypeCode02     string `json:"purchase_order_type_code_02,omitempty"`
	PurchaseOrderNumber03       string `json:"purchase_order_number_03,omitempty"`
	ReleaseNumber04             string `json:"release_number_04,omitempty"`
	Date05                      string `json:"date_05,omitempty"`
	ContractNumber06            string `json:"contract_number_06,omitempty"`
	AcknowledgmentType07        string `json:"acknowledgment_type_07,omitempty"`
	InvoiceTypeCode08           string `json:"invoice_type_code_08,omitempty"`
	ContractTypeCode09          string `json:"contract_type_code_09,omitempty"`
	PurchaseCategory10          string `json:"purchase_category_10,omitempty"`
	SecurityLevelCode11         string `json:"security_level_code_11,omitempty"`
	TransactionTypeCode12       string `json:"transaction_type_code_12,omitempty"`
}

type NameN1 struct {
	EntityIdentifierCode01        string `json:"entity_identifier_code_01,omitempty"`
	Name02                        string `json:"name_02,omitempty"`
	IdentificationCodeQualifier03 string `json:"identification_code_qualifier_03,omitempty"`
	IdentificationCode04          string `json:"identification_code_04,omitempty"`
	EntityRelationshipCode05      string `json:"entity_relationship_code_05,omitempty"`
	EntityIdentifierCode06        string `json:"entity_identifier_code_06,omitempty"`
}
type AdditionalNameInformationN2 struct {
	Name01 string `json:"name_01,omitempty"`
	Name02 string `json:"name_02,omitempty"`
}
type AddressInformationN3 struct {
	AddressInformation01 string `json:"address_information_01,omitempty"`
	AddressInformation02 string `json:"address_information_02,omitempty"`
}
type GeographicLocationN4 struct {
	CityName01            string `json:"city_name_01,omitempty"`
	StateOrProvinceCode02 string `json:"state_or_province_code_02,omitempty"`
	PostalCode03          string `json:"postal_code_03,omitempty"`
	CountryCode04         string `json:"country_code_04,omitempty"`
	LocationQualifier05   string `json:"location_qualifier_05,omitempty"`
	LocationIdentifier06  string `json:"location_identifier_06,omitempty"`
}
type LocationIDComponentNX2 struct {
	AddressComponentQualifier01 string `json:"address_component_qualifier_01,omitempty"`
	AddressInformation02        string `json:"address_information_02,omitempty"`
	CountyDesignator03          string `json:"county_designator_03,omitempty"`
}

type NameN1Loop struct {
	NameN1                      NameN1                        `json:"name_N1,omitempty"`
	AdditionalNameInformationN2 []AdditionalNameInformationN2 `json:"additional_name_information_N2,omitempty"`
	AddressInformationN3        []AddressInformationN3        `json:"address_information_N3,omitempty"`
	GeographicLocationN4        []GeographicLocationN4        `json:"geographic_location_N4,omitempty"`
	LocationIDComponentNX2      []LocationIDComponentNX2      `json:"location_id_component_NX2,omitempty"`
}
type NameN1Loop_Detail struct {
	NameN1                      NameN1                        `json:"name_N1,omitempty"`
	AdditionalNameInformationN2 []AdditionalNameInformationN2 `json:"additional_name_information_N2,omitempty"`
	AddressInformationN3        []AddressInformationN3        `json:"address_information_N3,omitempty"`
	GeographicLocationN4        *GeographicLocationN4         `json:"geographic_location_N4,omitempty"`
	LocationIDComponentNX2      []LocationIDComponentNX2      `json:"location_id_component_NX2,omitempty"`
}
type Heading struct {
	TransactionSetHeaderST              TransactionSetHeaderST              `json:"transaction_set_header_ST,omitempty"`
	BeginningSegmentForPurchaseOrderBEG BeginningSegmentForPurchaseOrderBEG `json:"beginning_segment_for_purchase_order_BEG,omitempty"`
	NameN1Loop                          []NameN1Loop                        `json:"name_N1_loop,omitempty"`
}
type BaselineItemDataPO1 struct {
	AssignedIdentification01        string  `json:"assigned_identification_01,omitempty"`
	QuantityOrdered02               float64 `json:"quantity_ordered_02"`
	UnitOrBasisForMeasurementCode03 string  `json:"unit_or_basis_for_measurement_code_03,omitempty"`
	UnitPrice04                     float64 `json:"unit_price_04"`
	BasisOfUnitPriceCode05          string  `json:"basis_of_unit_price_code_05,omitempty"`
	ProductServiceIDQualifier06     string  `json:"product_service_id_qualifier_06,omitempty"`
	ProductServiceID07              string  `json:"product_service_id_07,omitempty"`
	ProductServiceIDQualifier08     string  `json:"product_service_id_qualifier_08,omitempty"`
	ProductServiceID09              string  `json:"product_service_id_09,omitempty"`
	ProductServiceIDQualifier10     string  `json:"product_service_id_qualifier_10,omitempty"`
	ProductServiceID11              string  `json:"product_service_id_11,omitempty"`
	ProductServiceIDQualifier12     string  `json:"product_service_id_qualifier_12,omitempty"`
	ProductServiceID13              string  `json:"product_service_id_13,omitempty"`
	ProductServiceIDQualifier14     string  `json:"product_service_id_qualifier_14,omitempty"`
	ProductServiceID15              string  `json:"product_service_id_15,omitempty"`
	ProductServiceIDQualifier16     string  `json:"product_service_id_qualifier_16,omitempty"`
	ProductServiceID17              string  `json:"product_service_id_17,omitempty"`
	ProductServiceIDQualifier18     string  `json:"product_service_id_qualifier_18,omitempty"`
	ProductServiceID19              string  `json:"product_service_id_19,omitempty"`
	ProductServiceIDQualifier20     string  `json:"product_service_id_qualifier_20,omitempty"`
	ProductServiceID21              string  `json:"product_service_id_21,omitempty"`
	ProductServiceIDQualifier22     string  `json:"product_service_id_qualifier_22,omitempty"`
	ProductServiceID23              string  `json:"product_service_id_23,omitempty"`
	ProductServiceIDQualifier24     string  `json:"product_service_id_qualifier_24,omitempty"`
	ProductServiceID25              string  `json:"product_service_id_25,omitempty"`
}

type ProductItemDescriptionPID struct {
	ItemDescriptionType01              string `json:"item_description_type_01,omitempty"`
	ProductProcessCharacteristicCode02 string `json:"product_process_characteristic_code_02,omitempty"`
	AgencyQualifierCode03              string `json:"agency_qualifier_code_03,omitempty"`
	ProductDescriptionCode04           string `json:"product_description_code_04,omitempty"`
	Description05                      string `json:"description_05,omitempty"`
	SurfaceLayerPositionCode06         string `json:"surface_layer_position_code_06,omitempty"`
	SourceSubqualifier07               string `json:"source_subqualifier_07,omitempty"`
	YesNoConditionOrResponseCode08     string `json:"yes_no_condition_or_response_code_08,omitempty"`
	LanguageCode09                     string `json:"language_code_09,omitempty"`
}
type ProductItemDescriptionPIDLoop struct {
	ProductItemDescriptionPID ProductItemDescriptionPID `json:"product_item_description_PID,omitempty"`
}
type DestinationQuantitySDQ struct {
	UnitOrBasisForMeasurementCode01 string  `json:"unit_or_basis_for_measurement_code_01,omitempty"`
	IdentificationCodeQualifier02   string  `json:"identification_code_qualifier_02,omitempty"`
	IdentificationCode03            string  `json:"identification_code_03,omitempty"`
	Quantity04                      float64 `json:"quantity_04,omitempty"`
	IdentificationCode05            string  `json:"identification_code_05,omitempty"`
	Quantity06                      float64 `json:"quantity_06,omitempty"`
	IdentificationCode07            string  `json:"identification_code_07,omitempty"`
	Quantity08                      float64 `json:"quantity_08,omitempty"`
	IdentificationCode09            string  `json:"identification_code_09,omitempty"`
	Quantity10                      float64 `json:"quantity_10,omitempty"`
	IdentificationCode11            string  `json:"identification_code_11,omitempty"`
	Quantity12                      float64 `json:"quantity_12,omitempty"`
	IdentificationCode13            string  `json:"identification_code_13,omitempty"`
	Quantity14                      float64 `json:"quantity_14,omitempty"`
	IdentificationCode15            string  `json:"identification_code_15,omitempty"`
	Quantity16                      float64 `json:"quantity_16,omitempty"`
	IdentificationCode17            string  `json:"identification_code_17,omitempty"`
	Quantity18                      float64 `json:"quantity_18,omitempty"`
	IdentificationCode19            string  `json:"identification_code_19,omitempty"`
	Quantity20                      float64 `json:"quantity_20,omitempty"`
	IdentificationCode21            string  `json:"identification_code_21,omitempty"`
	Quantity22                      float64 `json:"quantity_22,omitempty"`
	LocationIdentifier23            string  `json:"location_identifier_23,omitempty"`
}

type BaselineItemDataPO1Loop struct {
	BaselineItemDataPO1           BaselineItemDataPO1             `json:"baseline_item_data_PO1,omitempty"`
	ProductItemDescriptionPIDLoop []ProductItemDescriptionPIDLoop `json:"product_item_description_PID_loop,omitempty"`
	DestinationQuantitySDQ        []DestinationQuantitySDQ        `json:"destination_quantity_SDQ,omitempty"`
	NameN1Loop                    []NameN1Loop_Detail             `json:"name_N1_loop,omitempty"`
}
type Detail struct {
	BaselineItemDataPO1Loop []BaselineItemDataPO1Loop `json:"baseline_item_data_PO1_loop,omitempty"`
}
type TransactionTotalsCTT struct {
	NumberOfLineItems01             int     `json:"number_of_line_items_01,omitempty"`
	HashTotal02                     float64 `json:"hash_total_02,omitempty"`
	Weight03                        float64 `json:"weight_03,omitempty"`
	UnitOrBasisForMeasurementCode04 string  `json:"unit_or_basis_for_measurement_code_04,omitempty"`
	Volume05                        float64 `json:"volume_05,omitempty"`
	UnitOrBasisForMeasurementCode06 string  `json:"unit_or_basis_for_measurement_code_06,omitempty"`
	Description07                   string  `json:"description_07,omitempty"`
}
type MonetaryAmountAMT struct {
	AmountQualifierCode01 string  `json:"amount_qualifier_code_01,omitempty"`
	MonetaryAmount02      float64 `json:"monetary_amount_02,omitempty"`
	CreditDebitFlagCode03 string  `json:"credit_debit_flag_code_03,omitempty"`
}
type TransactionTotalsCTTLoop struct {
	TransactionTotalsCTT *TransactionTotalsCTT `json:"transaction_totals_CTT,omitempty"`
	MonetaryAmountAMT    *MonetaryAmountAMT    `json:"monetary_amount_AMT,omitempty"`
}
type TransactionSetTrailerSE struct {
	NumberOfIncludedSegments01    int `json:"number_of_included_segments_01,omitempty"`
	TransactionSetControlNumber02 int `json:"transaction_set_control_number_02,omitempty"`
}
type Summary struct {
	TransactionTotalsCTTLoop []TransactionTotalsCTTLoop `json:"transaction_totals_CTT_loop,omitempty"`
	TransactionSetTrailerSE  TransactionSetTrailerSE    `json:"transaction_set_trailer_SE,omitempty"`
}

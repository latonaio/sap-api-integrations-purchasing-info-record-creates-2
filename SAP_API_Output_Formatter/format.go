package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	general := &General{
		PurchasingInfoRecord:        data.PurchasingInfoRecord,
		Supplier:                    data.Supplier,
		Material:                    data.Material,
		MaterialGroup:               data.MaterialGroup,
		PurgDocOrderQuantityUnit:    data.PurgDocOrderQuantityUnit,
		SupplierMaterialNumber:      data.SupplierMaterialNumber,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplierMaterialGroup:       data.SupplierMaterialGroup,
		IsRegularSupplier:           data.IsRegularSupplier,
		AvailabilityStartDate:       data.AvailabilityStartDate,
		AvailabilityEndDate:         data.AvailabilityEndDate,
		Manufacturer:                data.Manufacturer,
		CreationDate:                data.CreationDate,
		PurchasingInfoRecordDesc:    data.PurchasingInfoRecordDesc,
		LastChangeDateTime:          data.LastChangeDateTime,
		IsDeleted:                   data.IsDeleted,
	}

	return general, nil
}

func ConvertToPurchasingOrganizationPlant(raw []byte, l *logger.Logger) (*PurchasingOrganizationPlant, error) {
	pm := &responses.PurchasingOrganizationPlant{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchasingOrganizationPlant. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	purchasingOrganizationPlant := &PurchasingOrganizationPlant{
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		PurchasingInfoRecordCategory:   data.PurchasingInfoRecordCategory,
		PurchasingOrganization:         data.PurchasingOrganization,
		Plant:                          data.Plant,
		Supplier:                       data.Supplier,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurgDocOrderQuantityUnit:       data.PurgDocOrderQuantityUnit,
		PurchasingGroup:                data.PurchasingGroup,
		MinimumPurchaseOrderQuantity:   data.MinimumPurchaseOrderQuantity,
		StandardPurchaseOrderQuantity:  data.StandardPurchaseOrderQuantity,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		LastReferencingPurchaseOrder:   data.LastReferencingPurchaseOrder,
		LastReferencingPurOrderItem:    data.LastReferencingPurOrderItem,
		NetPriceAmount:                 data.NetPriceAmount,
		MaterialPriceUnitQty:           data.MaterialPriceUnitQty,
		PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
		PriceValidityEndDate:           data.PriceValidityEndDate,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		TaxCode:                        data.TaxCode,
		IncotermsClassification:        data.IncotermsClassification,
		MaximumOrderQuantity:           data.MaximumOrderQuantity,
		IsRelevantForAutomSrcg:         data.IsRelevantForAutomSrcg,
		IsEvaluatedRcptSettlmtAllowed:  data.IsEvaluatedRcptSettlmtAllowed,
		IsPurOrderAllwdForInbDeliv:     data.IsPurOrderAllwdForInbDeliv,
		IsOrderAcknRqd:                 data.IsOrderAcknRqd,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}

	return purchasingOrganizationPlant, nil
}

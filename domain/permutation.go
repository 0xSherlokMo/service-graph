package domain

import (
	"fmt"
	"strings"

	"github.com/graduation-fci/service-graph/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	EnglistNotificationDecreasedBy = "decreased by # interaction"
	EnglistNotificationIncreasedBy = "increased by # interaction"
	EnglishNotificationTemplate    = "Your medication interactions "
	ArabicNotificationDecreasedBy  = "قلت بعدد # تفاعل"
	ArabicNotificationIncreasedBy  = "زادت بعدد # تفاعل"
	ArabicNotificationTemplate     = "التفاعلات بين ادويتك "
)

type I18N struct {
	NameEn string `bson:"nameEn"`
	NameAr string `bson:"nameAr"`
}

func (i18n I18N) ToModel(i18nProto *proto.I18N) I18N {
	return I18N{
		NameEn: i18nProto.GetNameEn(),
		NameAr: i18nProto.GetNameAr(),
	}
}

type ReportInteraction struct {
	Drugs    []string `bson:"drugs"`
	Severity string   `bson:"severity"`
}

func (r ReportInteraction) ToModel(interaction *proto.Interaction) ReportInteraction {
	return ReportInteraction{
		Drugs:    interaction.GetDrugs(),
		Severity: interaction.GetSeverity(),
	}
}

type ReportPermutation struct {
	Medecines    []I18N              `bson:"medecines"`
	Interactions []ReportInteraction `bson:"interactions"`
}

func (r ReportPermutation) ToModel(permutation *proto.Permutation) ReportPermutation {
	var reportPermutation ReportPermutation
	for _, medecine := range permutation.Medecines {
		reportPermutation.Medecines = append(reportPermutation.Medecines, I18N{}.ToModel(medecine))
	}

	for _, interaction := range permutation.Interactions {
		reportPermutation.Interactions = append(reportPermutation.Interactions, ReportInteraction{}.ToModel(interaction))
	}

	return reportPermutation
}

type Report struct {
	XID                primitive.ObjectID  `bson:"_id"`
	MedicationId       int64               `bson:"medicationId"`
	ReportPermutations []ReportPermutation `bson:"reportPermutations"`
}

func (latestReport Report) BuildNotification(currentReport []*proto.Permutation) *proto.Notification {
	if len(latestReport.ReportPermutations) > len(currentReport) {
		arabicSuffix := strings.Replace(
			ArabicNotificationDecreasedBy,
			"#",
			fmt.Sprintf("%d", len(latestReport.ReportPermutations)-len(currentReport)),
			-1,
		)

		EnglishSuffix := strings.Replace(
			EnglistNotificationDecreasedBy,
			"#",
			fmt.Sprintf("%d", len(latestReport.ReportPermutations)-len(currentReport)),
			-1,
		)
		return &proto.Notification{
			Ar: ArabicNotificationTemplate + arabicSuffix,
			En: EnglishNotificationTemplate + EnglishSuffix,
		}
	}

	if len(latestReport.ReportPermutations) < len(currentReport) {
		arabicSuffix := strings.Replace(
			ArabicNotificationIncreasedBy,
			"#",
			fmt.Sprintf("%d", len(currentReport)-len(latestReport.ReportPermutations)),
			-1,
		)

		EnglishSuffix := strings.Replace(
			EnglistNotificationIncreasedBy,
			"#",
			fmt.Sprintf("%d", len(currentReport)-len(latestReport.ReportPermutations)),
			-1,
		)
		return &proto.Notification{
			Ar: ArabicNotificationTemplate + arabicSuffix,
			En: EnglishNotificationTemplate + EnglishSuffix,
		}
	}

	return &proto.Notification{}
}

func (r Report) ToModel(permutations []*proto.Permutation, medicationId int64) Report {
	var reportPermutations []ReportPermutation
	for _, permutation := range permutations {
		reportPermutations = append(reportPermutations, ReportPermutation{}.ToModel(permutation))
	}
	return Report{
		XID:                primitive.NewObjectID(),
		MedicationId:       medicationId,
		ReportPermutations: reportPermutations,
	}
}

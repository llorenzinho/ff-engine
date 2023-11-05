package fflag

import (
	models "ffapi/models/crud"
	"ffapi/pkg/apis/deployment/v1alpha1"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewFFfromContextBody(c *gin.Context) *v1alpha1.FeatureFlag {
	newFFmodel := &models.FeatureFlag{}
	if err := c.BindJSON(newFFmodel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}
	newFF := &v1alpha1.FeatureFlag{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newFFmodel.Name,
			Namespace: newFFmodel.Namespace,
		},
		Spec: v1alpha1.FeatureFlagSpec{
			Enabled: *newFFmodel.Enabled,
			Status:  *newFFmodel.Active,
		},
	}
	return newFF
}

func CastFFtoModel(ff *v1alpha1.FeatureFlag) *models.FeatureFlag {
	return &models.FeatureFlag{
		Name:      ff.Name,
		Namespace: ff.Namespace,
		Enabled:   &ff.Spec.Enabled,
		Active:    &ff.Spec.Status,
	}
}

func CastFFListtoModel(ffList *v1alpha1.FeatureFlagList) []*models.FeatureFlag {
	var modelList []*models.FeatureFlag
	for _, ff := range ffList.Items {
		modelList = append(modelList, CastFFtoModel(&ff))
	}
	return modelList
}

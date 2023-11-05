package fflag

// Gin handler for the featureflag resource

import (
	"encoding/json"
	"ffapi/pkg/apis/deployment/v1alpha1"
	k8s "ffapi/utils/k8sclient"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/api/errors"
)

func GetFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, CastFFtoModel(val))
}

func ListFeatureFlags(c *gin.Context) {
	namespace := c.Query("namespace")
	val, err := k8s.FFClient.List(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, CastFFListtoModel(val))
}

func CreateFeatureFlag(c *gin.Context) {
	newFF := NewFFfromContextBody(c)
	fmt.Println(newFF)
	val, err := k8s.FFClient.Create(newFF)
	if err != nil {
		if errors.IsAlreadyExists(err) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			c.AbortWithError(http.StatusConflict, err)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusCreated, CastFFtoModel(val))
}

func UpdateFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	input, _ := json.Marshal(c.Request.Body)
	newFF := &v1alpha1.FeatureFlag{}
	err := json.Unmarshal(input, newFF)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val, err := k8s.FFClient.Update(namespace, newFF)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, CastFFtoModel(val))
}

func DeleteFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	err := k8s.FFClient.Delete(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func IsFeatureFlagActive(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"enabled": val.Spec.Enabled, "active": val.Status.Active})
}

func EnableFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	val.Spec.Enabled = true
	_, err = k8s.FFClient.Update(namespace, val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func DisableFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	val.Spec.Enabled = false
	_, err = k8s.FFClient.Update(namespace, val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func ActivateFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if !val.Spec.Enabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "feature flag is not enabled"})
		return
	}
	val.Status.Active = true
	_, err = k8s.FFClient.UpdateStatus(namespace, val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func DeactivateFeatureFlag(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")
	val, err := k8s.FFClient.Get(namespace, name)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	val.Spec.Status = false
	_, err = k8s.FFClient.Update(namespace, val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Initialize the gin router
func InitFeatureFlagRoutes(r *gin.Engine) {
	// Those are referred to k8s resource itself
	r.GET("/api/v1/featureflags/:name", GetFeatureFlag)
	r.GET("/api/v1/featureflags", ListFeatureFlags)
	r.POST("/api/v1/featureflags", CreateFeatureFlag)
	r.PUT("/api/v1/featureflags", UpdateFeatureFlag)
	r.DELETE("/api/v1/featureflags/:name", DeleteFeatureFlag)

	// Those are referred to the status of the featureflag resource
	// Those endpoints must be used from applications that are using the featureflag resource
	r.GET("/api/v1/featureflags/:name/active", IsFeatureFlagActive)

	r.PUT("/api/v1/featureflags/:name/enable", EnableFeatureFlag)
	r.PUT("/api/v1/featureflags/:name/disable", DisableFeatureFlag)

	r.PUT("/api/v1/featureflags/:name/activate", ActivateFeatureFlag)
	r.PUT("/api/v1/featureflags/:name/deactivate", DeactivateFeatureFlag)
}

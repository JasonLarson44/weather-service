package weather

import (
	"testing"

	pb "github.com/jasonlarson44/weather-service/protobuf"
	"github.com/magiconair/properties/assert"
)

func TestTempToRating(t *testing.T) {
	temp := 25.0
	assert.Equal(t, "Cold", tempToRating(temp, pb.UnitsType_IMPERIAL.String()), "Temp should be Cold in imperial")
	assert.Equal(t, "Hot", tempToRating(temp, pb.UnitsType_METRIC.String()), "Temp should be Hot in Metric")
	assert.Equal(t, "Cold", tempToRating(temp, pb.UnitsType_STANDARD.String()), "Temp should be Cold in Kelvin")
}

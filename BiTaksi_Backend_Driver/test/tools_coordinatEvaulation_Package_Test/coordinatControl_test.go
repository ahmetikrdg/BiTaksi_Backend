package tools_coordinatEvaulation_Package_Test

import (
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/tools/coordinatEvaluation"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var drivers = []models.Coordinat{
	{
		Latitude:   41.01821618,
		Longtitude: 29.16796399,
	},
	{
		Latitude:   40.93057441,
		Longtitude: 29.0263735,
	},
	{
		Latitude:   40.98649307,
		Longtitude: 29.11370075,
	},
	{
		Latitude:   41.20368761,
		Longtitude: 29.10754687,
	},
	{
		Latitude:   41.06640224,
		Longtitude: 28.96196971,
	},
	{
		Latitude:   41.01821618,
		Longtitude: 29.16796399,
	},
	{
		Latitude:   41.10569174,
		Longtitude: 29.1061127,
	},
}
var coordinat = models.Coordinat{
	Latitude:   41.01721618,
	Longtitude: 29.16796399,
}

var coordinat2 = models.Coordinat{
	Latitude:   41.00721618,
	Longtitude: 29.16796399,
}

func TestCoordinatControl_Result_KM(t *testing.T) {

	result := coordinatEvaluation.CoordinatControl(coordinat, drivers)
	assert.NotNil(t, result)
	assert.NotEqual(t, result, 0)
	assert.EqualValues(t, result, 0.11119492664508966)
}

func TestCoordinatControl_Result_Count(t *testing.T) {

	result := coordinatEvaluation.CoordinatControl(coordinat, drivers)
	assert.NotNil(t, result)
	element := fmt.Sprintf("%.17f", result)
	count := len(element)
	assert.NotNil(t, count)
	assert.Equal(t, count, 19)
	assert.NotEqual(t, result, 0)
}

func TestCoordinatControl(t *testing.T) {
	type args struct {
		crd     models.Coordinat
		drivers []models.Coordinat
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Run", args{coordinat, drivers}, 0.11392664508966},
		{"Run2", args{coordinat2, drivers}, 1.2245193090456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.want, tt.name, tt.args)
			assert.NotEqual(t, tt.want, coordinatEvaluation.CoordinatControl(tt.args.crd, tt.args.drivers), "CoordinatControl(%v, %v)", tt.args.crd, tt.args.drivers)
		})
	}
}

func TestHaversine_Return_Value(t *testing.T) {
	var coordinat = models.Coordinat{
		Latitude:   41.01721618,
		Longtitude: 29.16796399,
	}

	result := coordinatEvaluation.Haversine(coordinat, 41.2144912, 29.09410704)
	assert.NotNil(t, result)
	assert.Equal(t, result, 22.79184258927576)
}

func TestHaversine_Return_Zero(t *testing.T) {
	var coordinat = models.Coordinat{
		Latitude:   0,
		Longtitude: 0,
	}

	result := coordinatEvaluation.Haversine(coordinat, 0, 0)
	assert.NotNil(t, result)
	assert.NotEqual(t, result, 22.79184258927576)
	assert.Equal(t, result, float64(0))
}

func TestHaversine_Multiple_Mismatched_Data_NotEqual(t *testing.T) {
	type args = struct {
		cr   models.Coordinat
		lat1 float64
		lon1 float64
	}
	var tests = []struct {
		name string
		args args
		want float64
	}{
		{"Haveresine wrong data", args{coordinat, 41.12084618, 29.07659061}, 14.836840955268126},
		{"haveresine wrong data 2", args{coordinat, 46.12084618, 22.07659061}, 14.836840955268126},
		{"haveresine wrong data 3", args{coordinat, 43.12084618, 29.07659061}, 14.836840955268126},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, args{})
			assert.NotEqual(t, tt.want, coordinatEvaluation.Haversine(tt.args.cr, tt.args.lat1, tt.args.lon1), "Haversine(%v, %v, %v)", tt.args.cr, tt.args.lat1, tt.args.lon1)
		})
	}
}

func TestHaversine_Multiple_Mismatched_Data_Equal(t *testing.T) {
	type args = struct {
		cr   models.Coordinat
		lat1 float64
		lon1 float64
	}
	var tests = []struct {
		name string
		args args
		want float64
	}{
		{"Haveresine Data 1", args{coordinat, 41.12084618, 29.07659061}, 13.836840955268126},
		{"haveresine Data 2", args{coordinat, 46.12084618, 22.07659061}, 804.6499878204428},
		{"haveresine Data 3", args{coordinat, 43.12084618, 29.07659061}, 234.03450349364},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, args{})
			assert.Equal(t, tt.want, coordinatEvaluation.Haversine(tt.args.cr, tt.args.lat1, tt.args.lon1), "Haversine(%v, %v, %v)", tt.args.cr, tt.args.lat1, tt.args.lon1)
		})
	}
}

func TestHaversine(t *testing.T) {
	mdls := models.Coordinat{Latitude: 0, Longtitude: 0}
	type args struct {
		cr   models.Coordinat
		lat1 float64
		lon1 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Haversine", args{mdls, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, coordinatEvaluation.Haversine(tt.args.cr, tt.args.lat1, tt.args.lon1), "Haversine(%v, %v, %v)", tt.args.cr, tt.args.lat1, tt.args.lon1)
		})
	}
}

func Test_kmCalculation_Multiple_Data(t *testing.T) {
	asm := []float64{5.1, 50, 60, 40, 80, 82, 85, 25, 54, 52}
	type args struct {
		km []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"KM Example", args{asm}, 5.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, coordinatEvaluation.KmCalculation(tt.args.km), "kmCalculation(%v)", tt.args.km)
		})
	}
}

func Test_kmCalculation(t *testing.T) {
	kmData := []float64{1.12084618, 29.07228447, 8.84269383, 1.08403379, 1.01403379, 0.08403379, 0.78403379, 0.78403379}
	result := coordinatEvaluation.KmCalculation(kmData)
	assert.NotNil(t, result)
	assert.Equal(t, result, 0.08403379)
	assert.IsType(t, result, 0.08403379)
	fmt.Println(result)
}

func Test_kmCalculation1(t *testing.T) {
	kmData := []float64{}
	panicF := func() {
		coordinatEvaluation.KmCalculation(kmData)
	}
	require.Panics(t, panicF)
}

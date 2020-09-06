package transformer

// func TestEmptyStringToJson(t *testing.T) {
// 	assert.Equal(t, "\"\"", support.NewValue("").ToJson())
// }
//
// func TestNilStringToJson(t *testing.T) {
// 	assert.Equal(t, "\"\"", support.NewValue(nil).ToJson())
// }
//
// func TestStringToJson(t *testing.T) {
// 	assert.Equal(t, "\"foo\"", support.NewValue("foo").ToJson())
// }
//
// func TestBoolToJson(t *testing.T) {
// 	assert.Equal(t, "true", support.NewValue(true).ToJson())
// 	assert.Equal(t, "false", support.NewValue(false).ToJson())
// }
//
// func TestNumberToJson(t *testing.T) {
// 	assert.Equal(t, "100", support.NewValue(100).ToJson())
// 	assert.Equal(t, "-100", support.NewValue(-100).ToJson())
// }
//
// func TestFloatToJson(t *testing.T) {
// 	assert.Equal(t, "0.1", support.NewValue(0.1).ToJson())
// }
//
// func TestCollectionWithOneStringToJson(t *testing.T) {
// 	assert.Equal(t, "[\"foo\"]", support.NewCollection("foo").ToJson())
// }
//
// func TestCollectionWithThoStringsToJson(t *testing.T) {
// 	assert.Equal(t, "[\"foo\",\"bar\"]", support.NewCollection("foo", "bar").ToJson())
// }
//
// func TestCollectionWithThoNumbersToJson(t *testing.T) {
// 	assert.Equal(t, "[12,14]", support.NewCollection(12, 14).ToJson())
// }
//
// func TestCollectionWithThoFloatToJson(t *testing.T) {
// 	assert.Equal(t, "[1.5,0.4]", support.NewCollection(1.5, 0.4).ToJson())
// }

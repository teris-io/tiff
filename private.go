package tiff

// TODO: Pass in the slice of valid FieldType for each tag.

// Private Tags >= 32768
// Reusable Tags >= 65000

// PrivateTags is exported to allow it to be modified directly with new tags.

var PrivateTags = NewTagSet("Private", 32768, 65535)

func init() {
	/*
		TODO:
		Most of these tags should actually be removed from this package
		and left as an exercise for users to register in their own
		packages. For example, if someone creates a GeoTIFF package, the
		Geo* tags would be defined in that package.  Those tags would be
		registered by importing that package.  We leave these tags here
		for reference and general use until such time as other proper
		packages represent them.
	*/

	// Do not lock PrivateTags.  This is meant to be added to by others.

	// http://www.awaresystems.be/imaging/tiff/tifftags/private.html
	PrivateTags.Register(NewTag(32932, "Wang Annotation", nil))
	PrivateTags.Register(NewTag(33445, "MD FileTag", nil))
	PrivateTags.Register(NewTag(33446, "MD ScalePixel", nil))
	PrivateTags.Register(NewTag(33447, "MD ColorTable", nil))
	PrivateTags.Register(NewTag(33448, "MD LabName", nil))
	PrivateTags.Register(NewTag(33449, "MD SampleInfo", nil))
	PrivateTags.Register(NewTag(33450, "MD PrepDate", nil))
	PrivateTags.Register(NewTag(33451, "MD PrepTime", nil))
	PrivateTags.Register(NewTag(33452, "MD FileUnits", nil))
	PrivateTags.Register(NewTag(33550, "ModelPixelScaleTag", nil))
	PrivateTags.Register(NewTag(33723, "IPTC", nil))
	PrivateTags.Register(NewTag(33918, "INGR Packet Data Tag", nil))
	PrivateTags.Register(NewTag(33919, "INGR Flag Registers", nil))
	PrivateTags.Register(NewTag(33920, "IrasB Transformation Matrix", nil))
	PrivateTags.Register(NewTag(33922, "ModelTiepointTag", nil))
	PrivateTags.Register(NewTag(34264, "ModelTransformationTag", nil))
	PrivateTags.Register(NewTag(34377, "Photoshop", nil))
	PrivateTags.Register(NewTag(34675, "ICC Profile", nil))
	PrivateTags.Register(NewTag(34735, "GeoKeyDirectoryTag", nil))
	PrivateTags.Register(NewTag(34736, "GeoDoubleParamsTag", nil))
	PrivateTags.Register(NewTag(34737, "GeoAsciiParamsTag", nil))
	PrivateTags.Register(NewTag(34908, "HylaFAX FaxRecvParams", nil))
	PrivateTags.Register(NewTag(34909, "HylaFAX FaxSubAddress", nil))
	PrivateTags.Register(NewTag(34910, "HylaFAX FaxRecvTime", nil))
	PrivateTags.Register(NewTag(37724, "ImageSourceData", nil))
	PrivateTags.Register(NewTag(40965, "Interoperability IFD", nil))
	PrivateTags.Register(NewTag(42112, "GDAL_METADATA", nil))
	PrivateTags.Register(NewTag(42113, "GDAL_NODATA", nil))
	PrivateTags.Register(NewTag(50215, "Oce Scanjob Description", nil))
	PrivateTags.Register(NewTag(50216, "Oce Application Selector", nil))
	PrivateTags.Register(NewTag(50217, "Oce Identification Number", nil))
	PrivateTags.Register(NewTag(50218, "Oce ImageLogic Characteristics", nil))
	PrivateTags.Register(NewTag(50706, "DNGVersion", nil))
	PrivateTags.Register(NewTag(50707, "DNGBackwardVersion", nil))
	PrivateTags.Register(NewTag(50708, "UniqueCameraModel", nil))
	PrivateTags.Register(NewTag(50709, "LocalizedCameraModel", nil))
	PrivateTags.Register(NewTag(50710, "CFAPlaneColor", nil))
	PrivateTags.Register(NewTag(50711, "CFALayout", nil))
	PrivateTags.Register(NewTag(50712, "LinearizationTable", nil))
	PrivateTags.Register(NewTag(50713, "BlackLevelRepeatDim", nil))
	PrivateTags.Register(NewTag(50714, "BlackLevel", nil))
	PrivateTags.Register(NewTag(50715, "BlackLevelDeltaH", nil))
	PrivateTags.Register(NewTag(50716, "BlackLevelDeltaV", nil))
	PrivateTags.Register(NewTag(50717, "WhiteLevel", nil))
	PrivateTags.Register(NewTag(50718, "DefaultScale", nil))
	PrivateTags.Register(NewTag(50719, "DefaultCropOrigin", nil))
	PrivateTags.Register(NewTag(50720, "DefaultCropSize", nil))
	PrivateTags.Register(NewTag(50721, "ColorMatrix1", nil))
	PrivateTags.Register(NewTag(50722, "ColorMatrix2", nil))
	PrivateTags.Register(NewTag(50723, "CameraCalibration1", nil))
	PrivateTags.Register(NewTag(50724, "CameraCalibration2", nil))
	PrivateTags.Register(NewTag(50725, "ReductionMatrix1", nil))
	PrivateTags.Register(NewTag(50726, "ReductionMatrix2", nil))
	PrivateTags.Register(NewTag(50727, "AnalogBalance", nil))
	PrivateTags.Register(NewTag(50728, "AsShotNeutral", nil))
	PrivateTags.Register(NewTag(50729, "AsShotWhiteXY", nil))
	PrivateTags.Register(NewTag(50730, "BaselineExposure", nil))
	PrivateTags.Register(NewTag(50731, "BaselineNoise", nil))
	PrivateTags.Register(NewTag(50732, "BaselineSharpness", nil))
	PrivateTags.Register(NewTag(50733, "BayerGreenSplit", nil))
	PrivateTags.Register(NewTag(50734, "LinearResponseLimit", nil))
	PrivateTags.Register(NewTag(50735, "CameraSerialNumber", nil))
	PrivateTags.Register(NewTag(50736, "LensInfo", nil))
	PrivateTags.Register(NewTag(50737, "ChromaBlurRadius", nil))
	PrivateTags.Register(NewTag(50738, "AntiAliasStrength", nil))
	PrivateTags.Register(NewTag(50740, "DNGPrivateData", nil))
	PrivateTags.Register(NewTag(50741, "MakerNoteSafety", nil))
	PrivateTags.Register(NewTag(50778, "CalibrationIlluminant1", nil))
	PrivateTags.Register(NewTag(50779, "CalibrationIlluminant2", nil))
	PrivateTags.Register(NewTag(50780, "BestQualityScale", nil))
	PrivateTags.Register(NewTag(50784, "Alias Layer Metadata", nil))

	DefaultTagSpace.RegisterTagSet(PrivateTags)
}

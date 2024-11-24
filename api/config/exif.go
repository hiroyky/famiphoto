package config

const ExifTagOrientation = 274
const ExifTagIDDateTimeOriginal = 36867

const (
	ExifOrientationNone                = 1 // 不要
	ExifOrientationHorizontal          = 2 // 水平方向に反転
	ExifOrientationRotate180           = 3 // 時計回りに180度回転
	ExifOrientationVertical            = 4 // 垂直方向に反転
	ExifOrientationHorizontalRotate270 = 5 // 水平方向に反転 + 時計回りに270度回転
	ExifOrientationRotate90            = 6 // 時計回りに90度回転
	ExifOrientationHorizontalRotate90  = 7 // 水平方向に反転 + 時計回りに90度回転
	ExifOrientationRotate270           = 8 // 時計回りに270度回転
)

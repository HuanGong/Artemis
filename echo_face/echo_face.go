package echo_face

import (
	"github.com/kelvins/lbph/metric"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/kelvins/lbph"
	"image"
	"os"
	"fmt"
)


type (
	EchoFace struct {
	}
)

func NewEchoFace() *EchoFace {
	logrus.SetLevel(logrus.DebugLevel)
	return &EchoFace{

	}
}

func (impl *EchoFace) BeforeCliRun() error {
	logrus.Debugln("EchoFace.BeforeCliRun Enter ")

	FaceData := []string {
		"gonghuan.png",
		"fxx.png",
		"cwj.png",
		"gs.png",
		"gh.png",
	}

	var images []image.Image
	for _, v := range FaceData {
		img, err := loadImage(v)
		checkError(err)
		images = append(images, img)
	}

	params := lbph.Params{
		Radius:    1,
		Neighbors: 4,
		GridX:     4,
		GridY:     4,
	}

	lbph.Init(params)

	err := lbph.Train(images, FaceData)
	checkError(err)

	logrus.Debugln("Echoface.BeforeCliRun Leave")
	return nil
}

func (impl *EchoFace) OnCliApplicationRun() error {
	return nil
}

func (impl *EchoFace) Endpoint() string {
	return "0.0.0.0:6667"
}
func (impl *EchoFace) HttpServerName() string {
	return "EchoFace"
}

func (impl *EchoFace) OnServerInitialized(ec *echo.Echo) error {
	logrus.Debugln("Server Initialize Enter")

	ec.GET("/test", func (ec echo.Context) error {

		logrus.Debugln("Handle Test Request Start Load Image")
		lbph.Metric = metric.EuclideanDistance

		img, err := loadImage("sgh.png")
		if err != nil {
			logrus.Debugln("Load Image sgh.png Failed", err.Error())
			return ec.String(200, "Load Image Failed")
		}

		logrus.Debugln("Handle Test Request Image Load Finish, => Start Predict")

		label, distance, err := lbph.Predict(img)
		if err != nil {
		    logrus.Debugln("Lable:", label, " Distance:", distance, " Error:", err.Error())
		}
		result := fmt.Sprintf("Lable:%s Distance:%f", label, distance)

		return ec.String(200, result)
	})

	impl.RunTest()
	return nil
}

func (impl *EchoFace) RunTest() {
	lbph.Metric = metric.EuclideanDistance

	img, err := loadImage("sfxx.png")
	if err != nil {
		logrus.Debugln("Load Image sfxx.png Failed", err.Error())
	}

	label, distance, err := lbph.Predict(img)
	if err != nil {
		logrus.Debugln("Lable:", label, " Distance:", distance, " Error:", err.Error())
	}
	logrus.Debugf("Lable:%s Distance:%f", label, distance)


	img, err = loadImage("sgh.png")
	if err != nil {
		logrus.Debugln("Load Image sfxx.png Failed", err.Error())
	}

	label, distance, err = lbph.Predict(img)
	if err != nil {
		logrus.Debugln("Lable:", label, " Distance:", distance, " Error:", err.Error())
	}
	logrus.Debugf("Lable:%s Distance:%f", label, distance)

	return
}


func loadImage(filePath string) (image.Image, error) {
	fImage, err := os.Open(filePath)
	checkError(err)

	defer fImage.Close()

	img, _, err := image.Decode(fImage)
	checkError(err)

	return img, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func LoadTrainModel() {

}


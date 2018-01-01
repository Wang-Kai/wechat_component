package lib

import (
	"testing"

	"github.com/goushuyun/log"
)

const (
	token  = "goushuyun"
	aesKey = "goushuyungoushuyungoushuyungoushuyungoushuy"
	appID  = "wx1c2695469ae47724"
)

func TestCrypt(t *testing.T) {
	text := `JpYHfqYGWgfuy/oO9QmbCCjVBT+HdOCEhvk98xhMqsKtAJg7k2/n1A287EWUc7HEubT/ReI+P2Alj8JeaCZqOHi858Cthiip/08fBu0pK4sPMHZ0NE2RvDw+hqVH9TfZ3pyK7ywHBxTZ3PkkmNrW3y/M/cS9MvLRdp5pZpK5tYlxskBDW2p8HrF/IFSlW11peBjcVDDEGI83wXrBnmxhFbx+LlexceTBGqoglcsfUdRWrMa222oOoPSk3xYF9TGHEH6Mu4YLLCqcA3PviquLQAOjbcm8ZFCrH8gKuijjLr5TUaEUGoFzm9e8tkkINqJ3s6ckNC6UaO++LPGtyR+PiPW1JPb6OYHE7JdPlUaGt+lkrs43zZbh0VuAI5q3zyNEpB2bRO8B+J4gtqYJYhNzyUXg2vU2io8BcbUKfDeYT2gQn90b+tk8e68oD8I9bjqXqwbh2NSQcXrdNFmrMJPHhg==`

	// decrypt
	crypter, err := NewmessageCrypter(token, aesKey, appID)
	if err != nil {
		log.Error(err)
	}

	crypterText, _, err := crypter.Decrypt(text)
	if err != nil {
		log.Error(err)
	}

	log.Debugf("the cryter xml is : %s", crypterText)
}

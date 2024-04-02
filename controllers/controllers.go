package controllers

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/Kagami/go-face"
	"github.com/swaggest/usecase"
)

func detect(rec *face.Recognizer) usecase.Interactor {
	type upload struct {
		Image multipart.File `formData:"image" description:"JPG image."`
	}

	type output struct {
		ElapsedSec float64     `json:"elapsedSec"`
		Found      int         `json:"found"`
		Faces      []face.Face `json:"faces,omitempty"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, in upload, out *output) (err error) {
		start := time.Now()
		imgData, err := io.ReadAll(in.Image)
		if err != nil {
			return err
		}

		out.Faces, err = rec.Recognize(imgData)
		out.Found = len(out.Faces)
		out.ElapsedSec = time.Since(start).Seconds()

		return err
	})

	u.SetTitle("Detect a face on image 'multipart/form-data'")

	return u
}

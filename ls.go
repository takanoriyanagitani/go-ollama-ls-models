package lmodels

import (
	"context"
	"encoding/json"

	oa "github.com/ollama/ollama/api"
)

type Oclient struct {
	*oa.Client
}

func EnvToClient() (Oclient, error) {
	cli, e := oa.ClientFromEnvironment()
	return Oclient{Client: cli}, e
}

func (o Oclient) List(ctx context.Context) (*oa.ListResponse, error) {
	return o.Client.List(ctx)
}

func (o Oclient) ListModels(ctx context.Context) ([]oa.ListModelResponse, error) {
	res, e := o.List(ctx)
	if nil != e {
		return nil, e
	}
	return res.Models, nil
}

func ResponseToJsonWriter(r *oa.ListModelResponse, w *json.Encoder) error {
	return w.Encode(r)
}

func ModelsToJsonWriter(models []oa.ListModelResponse, w *json.Encoder) error {
	for _, model := range models {
		if err := ResponseToJsonWriter(&model, w); err != nil {
			return err
		}
	}
	return nil
}

func (o Oclient) FullModelInfo(ctx context.Context, name string) (*oa.ShowResponse, error) {
	req := &oa.ShowRequest{Model: name}
	return o.Client.Show(ctx, req)
}

func FullModelInfoToJsonWriter(f *oa.ShowResponse, w *json.Encoder) error {
	return w.Encode(f)
}

func ModelName(r oa.ListModelResponse) string { return r.Model }

func (o Oclient) FullModelsToJsonWriter(
	ctx context.Context,
	models []oa.ListModelResponse,
	w *json.Encoder,
) error {
	for _, model := range models {
		var mname string = ModelName(model)
		fullInfo, err := o.FullModelInfo(ctx, mname)
		if err != nil {
			return err
		}

		err = FullModelInfoToJsonWriter(fullInfo, w)
		if err != nil {
			return err
		}
	}

	return nil
}

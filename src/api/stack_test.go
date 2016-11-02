package api

import (
	"testing"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/gin-gonic/gin"
)

func TestApi_UpdateStack(t *testing.T) {
	t.Skipf("TODO:not yet implemented.")
}

func TestApi_CreateStack(t *testing.T) {
	type fields struct {
		Client *dockerclient.CraneDockerClient
		Config *config.Config
	}
	type args struct {
		ctx *gin.Context
	}
	fakeClient := &dockerclient.CraneDockerClient{}
	fakeApi := &Api{
		Client: fakeClient,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test_create_stack",
			&dockerclient.CraneDockerClient{},
			&config.Config{},
			&gin.Context{},
		},
	}
	for _, tt := range tests {
		api := &Api{
			Client: tt.fields.Client,
			Config: tt.fields.Config,
		}
		api.CreateStack(tt.args.ctx)
	}
}

func TestApi_ListStack(t *testing.T) {
	type fields struct {
		Client *dockerclient.CraneDockerClient
		Config *config.Config
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		api := &Api{
			Client: tt.fields.Client,
			Config: tt.fields.Config,
		}
		api.ListStack(tt.args.ctx)
	}
}

func TestApi_InspectStack(t *testing.T) {
	type fields struct {
		Client *dockerclient.CraneDockerClient
		Config *config.Config
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		api := &Api{
			Client: tt.fields.Client,
			Config: tt.fields.Config,
		}
		api.InspectStack(tt.args.ctx)
	}
}

func TestApi_ListStackService(t *testing.T) {
	type fields struct {
		Client *dockerclient.CraneDockerClient
		Config *config.Config
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		api := &Api{
			Client: tt.fields.Client,
			Config: tt.fields.Config,
		}
		api.ListStackService(tt.args.ctx)
	}
}

func TestApi_RemoveStack(t *testing.T) {
	type fields struct {
		Client *dockerclient.CraneDockerClient
		Config *config.Config
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		api := &Api{
			Client: tt.fields.Client,
			Config: tt.fields.Config,
		}
		api.RemoveStack(tt.args.ctx)
	}
}

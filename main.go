package main

import (
	"context"
	"fmt"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/mittwald/goharbor-client/v5/apiv2"
)

func main() {
	fmt.Println("Starting The Use Cases")

	// Login To Harvour
	LoginTovtHarbor()

}

func LoginTovtHarbor() {
	client, err := apiv2.NewRESTClientForHost("http://localhost:8090/api", "admin", "godcracker123", nil)
	if err != nil {
		panic(err)
	}

	// projeRequest := model.ProjectReq{
	// 	ProjectName: "testi",
	// 	Metadata: &model.ProjectMetadata{
	// 		Public: "true",
	// 	},
	// }
	// err = client.NewProject(context.Background(), &projeRequest)

	// if err != nil {
	// 	panic(err)
	// }

	proj, err := client.GetProject(context.Background(), "tesli")
	//proj, err := client.ListProjects(context.Background(), "test")

	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(proj)
}

func LoginToHarvour() {
	clientSetConfig := &harbor.ClientSetConfig{
		URL:      "http://localhost:8090",
		Insecure: true,
		Username: "admin",
		Password: "godcracker123",
	}

	clientset, err := harbor.NewClientSet(clientSetConfig)
	clv2 := clientset.V2()
	getProject, err := clv2.Project.GetProject(context.Background(), &project.GetProjectParams{
		ProjectNameOrID: "test",
		Context:         context.TODO(),
	})
	if err != nil {
		fmt.Println("Error creating Harbor client:", err)
		return
	} else {
		fmt.Println(getProject.IsSuccess())
		fmt.Println(getProject.GetPayload().Name)
	}

	var (
		AutoScan                 = "false"
		EnableContentTrust       = "false"
		EnableContentTrustCosign = "false"
		PreventVul               = "false"
		Public                   = "false"
	)

	createProject, err := clv2.Project.CreateProject(context.Background(), &project.CreateProjectParams{
		Project: &models.ProjectReq{
			ProjectName: "test-client-v2",
			Metadata: &models.ProjectMetadata{
				AutoScan:                 &AutoScan,
				EnableContentTrust:       &EnableContentTrust,
				EnableContentTrustCosign: &EnableContentTrustCosign,
				PreventVul:               &PreventVul,
				Public:                   Public,
			},
		},
	})
	if err != nil {
		fmt.Println("Error creating Harbor client:", err)
		return
	} else {
		fmt.Println(createProject.IsSuccess())
		fmt.Println(createProject.String())
	}

}

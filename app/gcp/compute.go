package gcp

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"github.com/joho/godotenv"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/terraform"
	test_structure "github.com/tnnmuhandiram/terraform-gcp-poc/modules/test_structure"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/random"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("file not found")
	}
}

type Compute struct {
	ID          string
	ProjectName string
	InstaneID   string
	Zone        string
}

func ComputeCreate(w http.ResponseWriter, r *http.Request) {
	print("-=======================")

	// ctx := context.Background()
	// projectID := "postgress-cluster"

	// client, err := datastore.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// u1 := uuid.Must(uuid.NewV4())

	// kind := "Task"
	// name := u1.String()
	// taskKey := datastore.NameKey(kind, name, nil)

	// task := Gcp{
	// 	ID:          u1.String(),
	// 	ProjectName: os.Getenv("GCP_PROJECT"),
	// }

	// if _, err := client.Put(ctx, taskKey, &task); err != nil {
	// 	log.Fatalf("Failed to save task: %v", err)
	// }
	exampleDir := test_structure.CopyTerraformFolderToTemp("../", "examples/terraform-gcp-example")

	projectId := "postgress-cluster"
	zone := "us-east1-b"
	bucketName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))
	instanceName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))
	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"gcp_project_id": projectId,
			"zone":           zone,
			"instance_name":  instanceName,
			"bucket_name":    bucketName,
		},
	}

	out := terraform.InitAndApply(terraformOptions)
	fmt.Print(out)
}

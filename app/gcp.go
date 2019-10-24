package app

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"github.com/joho/godotenv"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/terraform"
	test_structure "github.com/tnnmuhandiram/terraform-gcp-poc/modules/test-structure"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/random"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("file not found")
	}
}

type Gcp struct {
	ID          string
	ProjectName string
}

func computeCreate(w http.ResponseWriter, r *http.Request) {

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

	// Get the Project Id to use
	projectId := "postgress-cluster"

	// Create all resources in the following zone
	zone := "us-east1-b"

	// Give the example bucket a unique name so we can distinguish it from any other bucket in your GCP account
	expectedBucketName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))

	// Also give the example instance a unique name
	expectedInstanceName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"gcp_project_id": projectId,
			"zone":           zone,
			"instance_name":  expectedInstanceName,
			"bucket_name":    expectedBucketName,
		},
	}

	out := terraform.InitAndApply(terraformOptions)
	fmt.Print(out)
}

func Deploy(w http.ResponseWriter, r *http.Request) {

}

// func Show(w http.ResponseWriter, r *http.Request) {

// 	db = config.DBConnection()
// 	user := User{}
// 	db.Where("id = ?", mux.Vars(r)["id"]).Find(&user)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }

// func AllUsers(w http.ResponseWriter, r *http.Request) {

// 	db = config.DBConnection()
// 	users := []User{}
// 	db.Find(&users)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }

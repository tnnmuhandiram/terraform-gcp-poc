package gcp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"

	structure "github.com/tnnmuhandiram/terraform-gcp-poc/modules/structure"
	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/terraform"
)

var token = "s.nyAMiGNtZbbFUyD98Ac9PpqJ"
var vault_addr = "http://127.0.0.1:8200"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("file not found")
	}
}

type Compute struct {
	ID        string
	ProjectID string
	InstaneID string
	Zone      string
	PublicIP  string
	BucketURL string
}

func ComputeCreate(w http.ResponseWriter, r *http.Request) {

	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	// secret, err := c.Read("secret/data/hello")
	secret, err := c.Read("secret/data/gcp/cred.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := json.Marshal(secret.Data)
	fmt.Println(string(b))

	// templateDir := structure.CopyTerraformFolderToTemp(".../", "scripts/compute-engine")
	// projectId := r.FormValue("project_id")
	// zone := r.FormValue("zone")
	// machineType := r.FormValue("machine_type")
	// bucketName := fmt.Sprintf("projectx-gcp-bucket-%s", strings.ToLower(random.UniqueId()))
	// instanceName := fmt.Sprintf("projectx-gcp-instance-%s", strings.ToLower(random.UniqueId()))
	// terraformOptions := &terraform.Options{
	// 	TerraformDir: templateDir,
	// 	Vars: map[string]interface{}{
	// 		"gcp_project_id": projectId,
	// 		"zone":           zone,
	// 		"instance_name":  instanceName,
	// 		"bucket_name":    bucketName,
	// 		"machine_type":   machineType,
	// 		"credentails_json" :
	// 	},
	// }
	// // defer terraform.Destroy(terraformOptions)
	// terraform.InitAndApply(terraformOptions)
	// // terraform.ApplyE(terraformOptions)
	// bucketURL := terraform.Output(terraformOptions, "bucket_url")
	// instanceOutputName := terraform.Output(terraformOptions, "instance_id")
	// publicIP := terraform.Output(terraformOptions, "public_ip")
	// // fmt.printf("%v", bucketURL)
	// fmt.Printf(instanceOutputName)
	// fmt.Printf(publicIP)

	// // fmt.Print(out)
	// ctx := context.Background()
	// // projectID := "postgress-cluster"

	// client, err := datastore.NewClient(ctx, projectId)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// u1 := uuid.Must(uuid.NewV4())

	// kind := "Terraform"
	// name := u1.String()
	// terraformKey := datastore.NameKey(kind, name, nil)

	// terraformData := Compute{
	// 	ID:        u1.String(),
	// 	ProjectID: projectId,
	// 	InstaneID: instanceOutputName,
	// 	Zone:      zone,
	// 	PublicIP:  publicIP,
	// 	BucketURL: bucketURL,
	// }

	// if _, err := client.Put(ctx, terraformKey, &terraformData); err != nil {
	// 	log.Fatalf("Failed to save terraform Data: %v", err)
	// }
}

func ComputeDestroy(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
	// kind := "Task"
	// ctx := context.Background()
	// client, err := datastore.NewClient(ctx, r.FormValue("project_id"))
	// print(err)
	// q := datastore.NewQuery("Task").Filter("uuid =", "gopher").Limit(1)

	// terraformKey := datastore.NameKey(kind, id, nil)
	// data := client.Get(ctx, terraformKey, nil)
	// // print(er)
	// fmt.Print(data)
	// fmt.Print(terraformKey)
	templateDir := structure.CopyTerraformFolderToTemp(".../", "scripts/compute-engine")
	terraformOptions := &terraform.Options{
		TerraformDir: templateDir,
		Vars: map[string]interface{}{
			"gcp_project_id": r.FormValue("project_id"),
			"zone":           r.FormValue("zone"),
			"instance_name":  r.FormValue("instance_name"),
			"bucket_name":    r.FormValue("bucket_name"),
			"machine_type":   r.FormValue("machine_type"),
		},
	}
	terraform.Destroy(terraformOptions)

}

package main

import (

	"encoding/json"

	"github.com/jxu86/kubeengine/core/common/log"
	"github.com/jxu86/kubeengine/core/common/util"
	"github.com/jxu86/kubeengine/core/kubeclient"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var logger = log.GetLogger("main", log.INFO)

func int32Ptr(i int32) *int32 { return &i }
func main() {
	// service.Server()
	logger.Infof("kube start")
	kc := kubeclient.NewClients("./config/kubeconfig.yaml")

	// kc.GetNodeList(metav1.ListOptions{})
	// kc.GetPodList("", metav1.ListOptions{})
	kc.GetNamespaceList(metav1.ListOptions{})

	contentBytes := util.Yamls2Bytes("./", []string{"./config/nginx-deployment.yaml"})
	jsonArray := util.Yamls2Jsons(contentBytes)
	deployment := &appsv1.Deployment{}
	for _, jsonObj := range jsonArray {
		// jsonObj := util.Yamls2Jsons(contentBytes)

		// deployment := &appsv1.Deployment{}
		err := json.Unmarshal(jsonObj, &deployment)
		if err != nil {
			logger.Error(err.Error())
		}
		// logger.Infof("deployment======>", deployment)
		break
	}

	// logger.Infof("contentBytes======>", contentBytes)
	// logger.Infof("jsonObj======>", string(jsonObj))

	// deployment := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "nginx-deployment",
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Replicas: int32Ptr(2),
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{
	// 				"app": "nginx",
	// 			},
	// 		},
	// 		Template: apiv1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Labels: map[string]string{
	// 					"app": "nginx",
	// 				},
	// 			},
	// 			Spec: apiv1.PodSpec{
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:  "nginx",
	// 						Image: "nginx:1.14.2",
	// 						Ports: []apiv1.ContainerPort{
	// 							{
	// 								Name:          "http",
	// 								Protocol:      apiv1.ProtocolTCP,
	// 								ContainerPort: 80,
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	logger.Infof("deployment======>", deployment)
	kc.CreateDeployment(deployment, metav1.CreateOptions{})
	kc.GetDeployment(deployment, metav1.GetOptions{})
	// kc.DeleteDeployment(deployment, metav1.DeleteOptions{})
	// kc.GetDeploymentList(deployment, metav1.ListOptions{})

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	// name := "client-go-test"
	// namespacesClient := clientset.CoreV1().Namespaces()
	// namespace := &apiv1.Namespace{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: name,
	// 	},
	// 	Status: apiv1.NamespaceStatus{
	// 		Phase: apiv1.NamespaceActive,
	// 	},
	// }
	// kc.DeleteNameSpace(name, metav1.DeleteOptions{})

	// kc.CreateNameSpace(name, metav1.CreateOptions{})

	// kc.GetNamespaceList(metav1.ListOptions{})
	// ns, err := kc.KubeClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	logger.Errorf(err.Error())
	// }
	// for _, n := range ns.Items {
	// 	logger.Infof("Node：", n.Name, n.Status.Addresses)
	// }

	// pods, err := kc.KubeClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	// logger.Infof("ns=>", ns)
	logger.Infof("end ...")

}

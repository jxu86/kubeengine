package kubeclient

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) GetNamespaceList(ops metav1.ListOptions) *corev1.NamespaceList {

	nss, err := c.KubeClient.CoreV1().Namespaces().List(context.TODO(), ops)
	if err != nil {
		logger.Errorf(err.Error())
	}
	for _, ns := range nss.Items {
		logger.Infof("Namespace：", ns.Name, ns.Status.Phase)
	}
	return nss
}

func (c *Clients) CreateNameSpace(name string, ops metav1.CreateOptions) *corev1.Namespace {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: corev1.NamespaceStatus{
			Phase: corev1.NamespaceActive,
		},
	}
	nameSpace, err := c.KubeClient.CoreV1().Namespaces().Create(context.TODO(), ns, ops)
	if err != nil {
		logger.Errorf(err.Error())
	}
	logger.Infof("Created namesapce %q \n", nameSpace.GetObjectMeta().GetName())
	return nameSpace
}

func (c *Clients) DeleteNameSpace(name string, ops metav1.DeleteOptions) error {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: corev1.NamespaceStatus{
			Phase: corev1.NamespaceActive,
		},
	}
	err := c.KubeClient.CoreV1().Namespaces().Delete(context.TODO(), name, ops)
	if err != nil {
		logger.Errorf(err.Error())
	}
	logger.Infof("Delete namesapce %q \n", ns.GetObjectMeta().GetName())
	return err
}

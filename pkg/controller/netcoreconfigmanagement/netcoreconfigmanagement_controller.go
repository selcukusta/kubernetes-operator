package netcoreconfigmanagement

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/types"

	coreerror "errors"

	selcukustav1alpha1 "github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_netcoreconfigmanagement")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NetCoreConfigManagement Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNetCoreConfigManagement{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {

	// Create a new controller
	c, err := controller.New("netcoreconfigmanagement-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	//Watch for changes to primary resource NetCoreConfigManagement
	err = c.Watch(&source.Kind{Type: &selcukustav1alpha1.NetCoreConfigManagement{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner NetCoreConfigManagement
	// Watch for configmap
	src := &source.Kind{Type: &corev1.ConfigMap{}}

	h := &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &selcukustav1alpha1.NetCoreConfigManagement{},
	}
	// h := &handler.EnqueueRequestForObject{}

	// fn := predicate.Funcs{
	// 	UpdateFunc: func(e event.UpdateEvent) bool {
	// 		log.Info(">>>>>  " + e.MetaNew.GetName())
	// 		oldConfigMap, oldConfigOk := e.ObjectOld.(*corev1.ConfigMap)
	// 		newConfigMap, newConfigOk := e.ObjectNew.(*corev1.ConfigMap)
	// 		if !oldConfigOk || !newConfigOk {
	// 			return false
	// 		}
	// 		different := !reflect.DeepEqual(oldConfigMap, newConfigMap)
	// 		log.Info(fmt.Sprintf("Configuration is changed: %v", different))
	// 		return different
	// 	},
	// }

	err = c.Watch(src, h /*fn*/)

	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNetCoreConfigManagement implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNetCoreConfigManagement{}

// ReconcileNetCoreConfigManagement reconciles a NetCoreConfigManagement object
type ReconcileNetCoreConfigManagement struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NetCoreConfigManagement object and makes changes based on the state read
// and what is in the NetCoreConfigManagement.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNetCoreConfigManagement) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NetCoreConfigManagement")
	instance := &selcukustav1alpha1.NetCoreConfigManagement{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		reqLogger.Error(err, err.Error())
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	currentConfigMap := &corev1.ConfigMap{}
	newConfigMap := r.createConfigMap(instance, map[string]string{instance.Spec.Config.ConfigMapKey: instance.Spec.Config.ConfigMapValue})
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Spec.Config.ConfigMapName, Namespace: request.Namespace}, currentConfigMap)
	if err == nil {
		if !reflect.DeepEqual(currentConfigMap.Data, newConfigMap.Data) {
			reqLogger.Info(fmt.Sprintf("ConfigMap will be changed: %v", instance.Spec.Config.ConfigMapName))
			if r.client.Update(context.TODO(), newConfigMap) != nil {
				return reconcile.Result{}, err
			}

			for _, deployment := range instance.Spec.LinkedDeployments {
				reqLogger.Info(fmt.Sprintf("Deployment will be restarted: %v", deployment))
				if r.restartDeployment(deployment, request.Namespace) != nil {
					return reconcile.Result{}, err
				}
			}
		}
	} else {
		if errors.IsNotFound(err) {
			if r.client.Create(context.TODO(), newConfigMap) != nil {
				reqLogger.Error(err, err.Error())
				return reconcile.Result{}, err
			}
		} else {
			reqLogger.Error(err, err.Error())
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileNetCoreConfigManagement) createConfigMap(cr *selcukustav1alpha1.NetCoreConfigManagement, data map[string]string) *corev1.ConfigMap {
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.Config.ConfigMapName,
			Namespace: cr.Namespace,
		},
		Data: data,
	}
	controllerutil.SetControllerReference(cr, cm, r.scheme)
	return cm
}

func (r *ReconcileNetCoreConfigManagement) restartDeployment(deploymentName string, deploymentNamespace string) error {
	currentDeployment := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: deploymentName, Namespace: deploymentNamespace}, currentDeployment)
	if err == nil {
		if currentDeployment.Spec.Template.ObjectMeta.Annotations == nil {
			currentDeployment.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		currentDeployment.Spec.Template.Annotations["kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		currentDeployment.Spec.Template.Annotations["kubernetes.io/change-cause"] = "Mounted config is changed."
		err = r.client.Update(context.TODO(), currentDeployment)
		if err != nil {
			return err
		}
		return nil
	}
	return coreerror.New("Unhandled exception")
}

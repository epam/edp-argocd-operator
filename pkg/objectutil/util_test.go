package objectutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	argoApi "github.com/epam/edp-argocd-operator/api/v1alpha1"
)

func TestContainsLabel(t *testing.T) {
	type args struct {
		object   v1.Object
		label    string
		labelVal string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "get labels success",
			args: args{
				object: &argoApi.ArgoApplication{
					ObjectMeta: v1.ObjectMeta{
						Labels: map[string]string{"label1": "val"},
					},
				},
				label:    "label1",
				labelVal: "val",
			},
			want: true,
		},
		{
			name: "empty labels",
			args: args{
				object: &argoApi.ArgoApplication{
					ObjectMeta: v1.ObjectMeta{
						Labels: nil,
					},
				},
				label:    "label2",
				labelVal: "val2",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsLabel(tt.args.object, tt.args.label, tt.args.labelVal); got != tt.want {
				t.Errorf("ContainsLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectMarkDeleted(t *testing.T) {
	type args struct {
		object client.Object
	}

	timeNow := v1.NewTime(time.Now())

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "object with deletion timestamp",
			args: args{
				object: &argoApi.ArgoApplication{
					ObjectMeta: v1.ObjectMeta{
						DeletionTimestamp: &timeNow,
					},
				},
			},
			want: true,
		},
		{
			name: "object with deletion timestamp and finalizer",
			args: args{
				object: &argoApi.ArgoApplication{
					ObjectMeta: v1.ObjectMeta{
						DeletionTimestamp: &timeNow,
						Finalizers:        []string{"finalizer"},
					},
				},
			},
			want: false,
		},
		{
			name: "object with finalizer",
			args: args{
				object: &argoApi.ArgoApplication{
					ObjectMeta: v1.ObjectMeta{
						Finalizers: []string{"finalizer"},
					},
				},
			},
			want: false,
		},
		{
			name: "no finalizers and deletion timestamp",
			args: args{
				object: &argoApi.ArgoApplication{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ObjectMarkDeleted(tt.args.object)

			assert.Equal(t, tt.want, got)
		})
	}
}

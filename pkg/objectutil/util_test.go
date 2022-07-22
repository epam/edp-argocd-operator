package objectutil

import (
	"github.com/epam/edp-argocd-operator/api/v1alpha1"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
				object: &v1alpha1.ArgoApplication{
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
				object: &v1alpha1.ArgoApplication{
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

// More of a test to just get comfortable with validator.v8
package msgspec_test
import (
	"testing"
	"github.com/byrnedo/blogsvc/msgspec"
)

func TestRequiredMembersSupplied(t *testing.T){
	var post = msgspec.NewPost{
		ID : "ID",
		Post: msgspec.Post{
			Title: "Title",
		},
	}

	errs := post.Validate()
	if len(errs) !=0 {
		t.Errorf("Got validation errors: %v\n", errs)
	}
}

func TestRequiredMembersMissing(t *testing.T) {
	var post = msgspec.NewPost{
		Post: msgspec.Post{
			Title: "",
		},
	}
	errs := post.Validate()
	if len(errs) != 2 {
		t.Errorf("Should have gotten 2 validation errors, instead got %d\n", len(errs))
	}

	if errs["NewPost.ID"] == nil {
		t.Errorf("Should have error for ID field, %+v\n", errs)
	}

	if errs["NewPost.Post.Title"] == nil {
		t.Error("Should have error for Title field")
	}
}



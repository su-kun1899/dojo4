package imgconv_test

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo4/kadai1/su-kun1899/imgconv"
)

func TestPngConv_Convert(t *testing.T) {
	type args struct {
		src  string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "png to jpeg",
			args: args{
				src:  filepath.Join("testdata", "syokuji_computer.jpg"),
				dest: filepath.Join("testdata", "syokuji_computer.png"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			j := imgconv.PngConv{}
			defer func() {
				if err := os.Remove(tt.args.dest); err != nil {
					t.Error("unexpected error:", err)
					return
				}
			}()

			// when
			err := j.Convert(tt.args.src, tt.args.dest)

			// then
			if (err != nil) != tt.wantErr {
				t.Errorf("PngConv.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// and
			if format := getFormat(t, tt.args.dest); format != "png" {
				t.Errorf("format = %v, want %v", format, "png")
			}
		})
	}
}

func getFormat(t *testing.T, path string) string {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer file.Close()

	_, format, err := image.Decode(file)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return format
}

func TestReplaceExt(t *testing.T) {
	type args struct {
		fileName string
		newExt   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "txt to dat", args: args{fileName: "foo.txt", newExt: "dat"}, want: "foo.dat"},
		{name: "jpg to png", args: args{fileName: "foo.jpg", newExt: "png"}, want: "foo.png"},
		{
			name: "path included",
			args: args{
				fileName: filepath.Join("foo", "baz", "bar.jpg"),
				newExt:   "png",
			},
			want: filepath.Join("foo", "baz", "bar.png"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgconv.ReplaceExt(tt.args.fileName, tt.args.newExt); got != tt.want {
				t.Errorf("ReplaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

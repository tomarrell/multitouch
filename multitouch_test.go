package multitouch

import (
	"testing"
)

func TestMultitouch_single_touches(t *testing.T) {
	mt, err := NewMultitouch("./testdata/evdump")
	if err != nil {
		t.Fatal(err)
	}

	go mt.Begin()

	// First touch top left
	e := mt.Next()
	if e.Action != ActionBegin {
		t.Fatal("action not begin")
	}

	if e.X < 480*0.25 || e.X > 480*0.35 {
		t.Fatalf("expected X coord in top left, got: %v", e.X)
	} else if e.Y < 480*0.25 || e.Y > 480*0.35 {
		t.Fatalf("expected Y coord in top left, got: %v", e.Y)
	}

	e = mt.Next()
	if e.Action != ActionEnd {
		t.Fatal("action not end")
	}

	// Second touch bottom left
	e = mt.Next()
	if e.Action != ActionBegin {
		t.Fatal("action not begin")
	}

	e = mt.Next()
	if e.Action != ActionEnd {
		t.Fatal("action not end")
	}

	// Third touch bottom right
	e = mt.Next()
	if e.Action != ActionBegin {
		t.Fatal("action not begin")
	}

	e = mt.Next()
	if e.Action != ActionEnd {
		t.Fatal("action not end")
	}

	// Fourth touch top right
	e = mt.Next()
	if e.Action != ActionBegin {
		t.Fatal("action not begin")
	}

	e = mt.Next()
	if e.Action != ActionEnd {
		t.Fatal("action not end")
	}
}

func Test_transformPoint(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name   string
		args   args
		wantXp int
		wantYp int
	}{
		{"top left corner", args{153, 350}, 130, 153},
		{"bottom left corner", args{153, 350}, 130, 153},
		{"top right corner", args{31, 12}, 468, 31},
		{"bottom right corner", args{457, 20}, 460, 457},
		{"top left corner 2", args{20, 454}, 26, 20},
		{"bottom left corner 2", args{462, 461}, 19, 462},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXp, gotYp := transformPoint(tt.args.x, tt.args.y)
			if gotXp != tt.wantXp {
				t.Errorf("tansformPoint() gotXp = %v, want %v", gotXp, tt.wantXp)
			}

			if gotYp != tt.wantYp {
				t.Errorf("tansformPoint() gotYp = %v, want %v", gotYp, tt.wantYp)
			}
		})
	}
}

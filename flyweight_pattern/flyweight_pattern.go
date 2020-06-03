package flyweight_pattern

type TreeType struct {
	name    string
	color   string
	texture string
}

func NewTreeType(name, color, texture string) TreeType {
	return TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

func (t *TreeType) draw(canvas string, x, y int) {
	// 绘制
}

type TreeFactory struct {
	treeTypes []TreeType
}

func (tf *TreeFactory) GetTreeType(name, color, texture string) TreeType {
	for _, t := range tf.treeTypes {
		if t.name == name && t.color == color && t.texture == color {
			return t
		}
	}
	return TreeType{}
}

var factory TreeFactory

type Tree struct {
	x        int
	y        int
	treeType TreeType
}

func NewTree(x, y int, treeType TreeType) Tree {
	return Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

func (t *Tree) draw(canvas string) {
	t.treeType.draw(canvas, t.x, t.y)
}

type Forest struct {
	trees []Tree
}

func (f *Forest) PlantTrees(x, y int, name, color, texture string) {
	treeType := factory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw(canvas string) {
	for _, t := range f.trees {
		t.draw(canvas)
	}
}

package animal


type Animal interface {
    Eat() string
    Sound() string
    Move() string
    Age() int
}


type Lion struct {
    Age int
}

func (l *Lion) Sound() string {
    return "Лев Рычит"
}

func (l *Lion) Move() string {
    return "Лев шагает"
}

func (l *Lion) Agee() int {
    return l.Age
}

func (l *Lion) Eat() string {
    return "Других животных"
}


type Giraffe struct {
    Age int
}

func (g *Giraffe) Sound() string {
    return "Жираф не знаю"
}

func (g *Giraffe) Move() string {
    return "Жираф шагает"
}

func (g *Giraffe) Agee() int {
    return g.Age
}

func (g *Giraffe) Eat() string {
    return "Жираф ест листья с дерева"
}


type Snake struct {
    Age int
}

func (s *Snake) Sound() string {
    return "Змея Шипит"
}

func (s *Snake) Move() string {
    return "Змея ползет"
}

func (s *Snake) Agee() int {
    return s.Age
}

func (s *Snake) Eat() string {
    return "Ест мышей"
}

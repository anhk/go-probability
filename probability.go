package probability

type Option struct {
	Percent int
}

type Probability struct {
	Percent int // 当前的运行概率，单位是%，默认为100
}

func (p *Probability) Should() bool {
	return RandomInt(1, 100) <= p.Percent
}

func (p *Probability) SetPercent(percent int) {
	p.Percent = (&Probability{Percent: percent}).valid().Percent
}

func (p *Probability) valid() *Probability {
	p.Percent = If(p.Percent > 100, 100, p.Percent)
	p.Percent = If(p.Percent < 0, 0, p.Percent)
	return p
}

func NewProbability(o *Option) *Probability {
	return (&Probability{
		Percent: o.Percent,
	}).valid()
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

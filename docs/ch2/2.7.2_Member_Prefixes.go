package ch2

type m_Part struct {
    m_dsc string // The textual description
}

func (p *m_Part) setName(name string) {
    p.m_dsc = name
}

type Part struct {
    description string // The textual description
}

func (p *Part) setDescription(description string) {
    p.description = description
}

package carteira

import "fmt"

type Conta struct {
	saldoEmCentavos int64
}

func (c *Conta) Depositar(centavos int64) {
	c.saldoEmCentavos += centavos
}

func (c *Conta) Sacar(centavos int64) error {
	if centavos > c.saldoEmCentavos {
		return fmt.Errorf("saldo insuficiente: disponível R$ %d,%02d", c.saldoEmCentavos/100, c.saldoEmCentavos%100)
	}
	c.saldoEmCentavos -= centavos
	return nil
}

func (c *Conta) Saldo() int64 {
	return c.saldoEmCentavos
}

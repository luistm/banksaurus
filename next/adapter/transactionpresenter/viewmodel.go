package transactionpresenter

import (
	"fmt"
)

// ViewModel contains data to be shown
type ViewModel struct {
	raw []string

	// Receives a view... which must be an io.writer
}

func (vm *ViewModel) String() string {

	// Here parse data with tabe writer
	//if c.output == nil {
	//	return errOutputPipeUndefined
	//}
	//
	//const padding = 1
	//w := tabwriter.NewWriter(c.output, 0, 0, padding, ' ', 0)
	//
	//for _, entity := range entities {
	//	switch entity.(type) {
	//	case *transaction.Transaction:
	//		stringArray := strings.Fields(entity.String())
	//		value := stringArray[0]
	//		seller := strings.Join(stringArray[1:], " ")
	//		fmt.Fprintf(w, "%s\t%s\n", value, seller)
	//	default:
	//		fmt.Fprintf(w, "%s\n", entity)
	//	}
	//}
	//w.Flush()

	return fmt.Sprintf("%s", vm.raw)
}

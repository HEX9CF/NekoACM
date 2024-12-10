package cmd

import "bufio"

func clearBuffer(reader *bufio.Reader) error {
	for reader.Buffered() > 0 {
		err := clearBuffer(reader)
		if err != nil {
			return err
		}
	}
	return nil
}

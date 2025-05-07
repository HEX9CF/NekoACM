package cli

import "bufio"

// 清空缓冲区
func clearBuffer(reader *bufio.Reader) error {
	for reader.Buffered() > 0 {
		_, err := reader.Discard(reader.Buffered())
		if err != nil {
			return err
		}
	}
	return nil
}

package cmd

import "bufio"

// 清空缓冲区
func clearBuffer(reader *bufio.Reader) error {
	for reader.Buffered() > 0 {
		err := clearBuffer(reader)
		if err != nil {
			return err
		}
	}
	return nil
}

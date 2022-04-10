package vk

// SendByParams - метод отправк сообщений в соответствии с параметрами
func (c ConnectionImpl) SendByParams(params map[string]interface{}) error {
	if _, err := c.vk.MessagesSend(params); err != nil {
		return err
	}

	return nil
}

import React from 'react'

import { Button } from '.'

export default {
  component: Button,
  title: 'ui/Button',
}

const buttonStory = (args) => (
  <Button onClick={() => null} { ...args }>
    {'i\'m a Button'}
  </Button>
)

export const Default = buttonStory.bind({})

export const disabled = buttonStory.bind({})
disabled.args = {
  disabled: true,
}

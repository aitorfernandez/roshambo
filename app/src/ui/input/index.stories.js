import React from 'react'

import { Input } from '.'

export default {
  component: Input,
  title: 'ui/Input',
}

const InputStory = (args) => (
  <Input { ...args } />
)

export const Default = InputStory.bind({})

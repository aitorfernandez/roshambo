import React from 'react'

import { Heading } from '.'

export default {
  title: 'ui/Heading',
  component: Heading,
}

const headingStory = (args) => (
  <Heading { ...args }>
    {'I\'m a Heading'}
  </Heading>
)

export const Default = headingStory.bind({})

export const levelTwo = headingStory.bind({})
levelTwo.args = {
  level: 2,
}

import React from 'react'

import { Flex } from '.'

export default {
  component: Flex,
  title: 'ui/Flex',
}

const BoxStory = (args) => (
  <Flex { ...args }>
    <div>
      Top
    </div>
    <div>
      Bottom
    </div>
  </Flex>
)

export const Default = BoxStory.bind({})

export const fontSize = BoxStory.bind({})
fontSize.args = {
  fs: 2,
}

export const column = BoxStory.bind({})
column.args = {
  direction: 'column',
}

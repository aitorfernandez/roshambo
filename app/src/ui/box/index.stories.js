import React from 'react'

import { Box } from '.'

export default {
  component: Box,
  title: 'ui/Box',
}

const BoxStory = (args) => (
  <Box { ...args }>
    Div content
  </Box>
)

export const fontSize = BoxStory.bind({})
fontSize.args = {
  fs: 2,
}

export const margin = BoxStory.bind({})
margin.args = {
  m: [5, 0, 5, 0],
}

export const padding = BoxStory.bind({})
padding.args = {
  p: [0, 5, 0, 5],
}

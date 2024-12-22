import React from 'react'

import { Paragraph } from '.'

export default {
  component: Paragraph,
  title: 'ui/Paragraph',
}

const ParagraphStory = (args) => (
  <Paragraph { ...args }>
    Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
  </Paragraph>
)

export const Default = ParagraphStory.bind({})

export const textCenter = ParagraphStory.bind({})
textCenter.args = {
  ta: 'center',
}

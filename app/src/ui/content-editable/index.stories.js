import React from 'react'

import { ContentEditable } from '.'

export default {
  component: ContentEditable,
  title: 'ui/ContentEditable',
}

const contentEditableStory = (args) => (
  <ContentEditable
    { ...args }
    onEdit={() => console.log('onEdit')}
  />
)

export const Default = contentEditableStory.bind({})

export const value = contentEditableStory.bind({})
value.args = {
  value: 'username',
}

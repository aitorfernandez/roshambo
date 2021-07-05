import React from 'react'

import { Icon } from '.'
import { theme } from '../../theme'

export default {
  component: Icon,
  title: 'ui/Icon',
}

const iconStory = (args) => <Icon { ...{ ...args, theme } } />

export const handPaper = iconStory.bind({})
handPaper.args = {
  icon: 'hand-paper',
}

export const handRock = iconStory.bind({})
handRock.args = {
  icon: 'hand-rock',
}

export const handScissors = iconStory.bind({})
handScissors.args = {
  icon: 'hand-scissors',
}

export const question = iconStory.bind({})
question.args = {
  icon: 'question',
}

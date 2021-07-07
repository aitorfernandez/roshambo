import React from 'react'

import { ButtonBase } from '.'

export default {
  title: 'ui/ButtonBase',
  component: ButtonBase,
}

function onClick() {
  console.log('onClick')
}

export const buttonBase = () => (
  <ButtonBase onClick={onClick}>
    {'I\'m a ButtonBase'}
  </ButtonBase>
)

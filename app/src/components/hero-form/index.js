import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { useForm } from 'react-hook-form'

import { t } from '../../theme'
import {
  Button,
  Flex,
  Heading,
  Input,
  Paragraph,
} from '../../ui'

export function HeroForm({
  getStarted,
  response,
}) {
  const {
    handleSubmit,
    register,
    formState: { isValid },
  } = useForm({ mode: 'onChange' })

  async function onSubmit(values) {
    await getStarted({
      variables: {
        ...values,
      },
    })
  }

  const form = (
    <Form
      onSubmit={handleSubmit(onSubmit)}
    >
      <StyledInput
        placeholder="Email here for start..."
        { ...register("email", { required: true, pattern: /^\S+@\S+$/i})}
      />
      <StyledButton
        disabled={!isValid}
        type="submit"
      >
        Get Started
      </StyledButton>
    </Form>
  )

  const message = (
    <Paragraph>
      Great! Check your email for the magic link.
    </Paragraph>
  )

  return (
    <Flex
      direction="column"
    >
      <Heading>
        Rock. Paper. Scissors.
      </Heading>
      {response ? message : form}
    </Flex>
  )
}

HeroForm.propTypes = {
  getStarted: PropTypes.func.isRequired,
  response: PropTypes.bool,
}

const Form = styled(Flex).attrs({
  as: 'form',
})`
  width: 100%;
`

const StyledInput = styled(Input)`
  flex: 4 0 0;
`

const StyledButton = styled(Button)`
  flex: 2 0 0;
  margin-left: ${t('space.4')};
`

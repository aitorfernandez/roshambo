import React from 'react'
import PropTypes from 'prop-types'
import { useForm } from 'react-hook-form'

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
    <form
      onSubmit={handleSubmit(onSubmit)}
    >
      <input
        placeholder="Email here for start..."
        { ...register("email", { required: true, pattern: /^\S+@\S+$/i})}
      />
      <button
        disabled={!isValid}
        type="submit"
      >
        Get Started
      </button>
    </form>
  )

  const message = (
    <p>
      Great! Check your email for the magic link.
    </p>
  )

  return (
    <div>
      <h2>
        Rock. Paper. Scissors.
      </h2>
      {response ? message : form}
    </div>
  )
}

HeroForm.propTypes = {
  getStarted: PropTypes.func.isRequired,
  response: PropTypes.bool,
}

import React from 'react'
import PropTypes from 'prop-types'
import { useCookies } from 'react-cookie'

export function StatCreateButton({
  createStat,
  playerMove,
}) {
  const [{ id: accountID }] = useCookies(['id'])

  async function onSubmit(values) {
    await createStat({
      variables: {
        input: {
          accountID,
          ...values,
        },
      },
    })
  }

  return (
    <div>
      <button
        onClick={() => onSubmit({ playerMove })}
      >
        StatCreateButton
      </button>
    </div>
  )
}

StatCreateButton.propTypes = {
  createStat: PropTypes.func.isRequired,
  playerMove: PropTypes.number.isRequired,
}

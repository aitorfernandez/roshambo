import React from 'react'
import {
  gql,
  useMutation,
  useQuery,
} from '@apollo/client'

import { useCookies } from 'react-cookie'
import { AppContent } from '../../components/app-content'
import { Header } from '../../components/header'
import { Profile } from '../../components/profile'
import { StatCreateButton } from '../../components/stat-create-button'
import { StatDeleteButton } from '../../components/stat-delete-button'
import { StatItem } from '../../components/stat-item'
import { StatResult } from '../../components/stat-result'
import { StatScore } from '../../components/stat-score'

export function PlayPage() {
  const [{ id }] = useCookies(['id'])

  const { data, loading, error } = useQuery(PlayPageAccountQuery, {
    variables: {
      id,
    },
  })

  const [createStat] = useMutation(PlayPageCreateStatMutation, {
    update: (cache, { data }) => {
      console.log(cache, data)
    },
  })
  const [deleteStats] = useMutation(PlayPageDeleteStatsMutation, {
    update: (cache) => {
      console.log(cache)
    },
  })

  if (loading) {
    return null
  }
  if (error) {
    console.error(error)
    return null
  }

  return (
    <AppContent
      header={
        <>
          <Header />
          <Profile />
        </>
      }
    >
      <StatResult { ...data.account } />
      <StatScore { ...data.account } />
      <div>
        <StatCreateButton { ...{ createStat } } />
        <StatCreateButton { ...{ createStat } } />
        <StatCreateButton { ...{ createStat } } />
        <StatDeleteButton { ...{ deleteStats } } />
      </div>
      <StatItem />
    </AppContent>
  )
}

const PlayPageAccountQuery = gql`
  query PlayPageAccountQuery(
    $id: ID!
  ) {
    account(id: $id) {
      id
    }
  }
`

const PlayPageCreateStatMutation = gql`
  mutation PlayPageCreateStatMutation(
    $input: CreateStatInput!
  ) {
    stat: createStat(input: $input) {
      id
    }
  }
`

const PlayPageDeleteStatsMutation = gql`
  mutation PlayPageDeleteStatsMutation(
    $accountID: ID!
  ) {
    deleteStatsByAccount(accountID: $accountID)
  }
`

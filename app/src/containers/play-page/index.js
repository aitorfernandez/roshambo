import React from 'react'
import {
  gql,
  useMutation,
  useQuery,
} from '@apollo/client'
import { useCookies } from 'react-cookie'

import { AppContent } from '../../components/app-content'
import { Flex } from '../../ui'
import { Footer } from '../../components/footer'
import { fromAccount } from '../../store'
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
      fromAccount.addStat(cache, data.stat, { id })
    },
  })
  const [deleteStats] = useMutation(PlayPageDeleteStatsMutation, {
    update: (cache) => {
      cache.reset()
    },
  })
  const [setProfile] = useMutation(PlayPageSetProfileMutation, {
    update: (cache, { data }) => {
      fromAccount.addProfile(cache, data.profile, {
        id,
      })
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
          <Profile { ...{ setProfile, accountID: data.account.id, ...data.account?.profile } } />
        </>
      }
      footer={<Footer />}
    >
      <StatResult { ...data.account.stats[0] } />
      <StatScore { ...data.account } />
      <Flex>
        <StatCreateButton { ...{ createStat, playerMove: 0 } } />
        <StatCreateButton { ...{ createStat, playerMove: 1 } } />
        <StatCreateButton { ...{ createStat, playerMove: 2 } } />
        <StatDeleteButton { ...{ deleteStats } } />
      </Flex>
      {data.account.stats.map((s) => (
        <StatItem key={s.id} { ...s } />
      ))}
    </AppContent>
  )
}

export const PlayPageAccountQuery = gql`
  query PlayPageAccountQuery(
    $id: ID!
  ) {
    account(id: $id) {
      id
      stats {
        ...StatItemStatFragment
      }
      profile {
        ...ProfileFragment
      }
    }
  }
  ${StatItem.fragments.stat}
  ${Profile.fragments.profile}
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

const PlayPageSetProfileMutation = gql`
  mutation PlayPageSetProfileMutation(
    $input: SetProfileInput!
  ) {
    profile: setProfile(input: $input) {
      ...ProfileFragment
    }
  }
  ${Profile.fragments.profile}
`

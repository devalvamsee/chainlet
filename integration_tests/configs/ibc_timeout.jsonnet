local ibc = import 'ibc.jsonnet';

ibc {
  'chainlet_988-1'+: {
    key_name: 'signer3',
    accounts: super.accounts + [{
      name: 'signer3',
      coins: '0clt',
      mnemonic: '${SIGNER3_MNEMONIC}',
    }],
    genesis+: {
      app_state+: {
        chainlet+: {
          params+: {
            ibc_timeout: 0,
          },
        },
      },
    },
  },
  relayer+: {
    chains: [super.chains[0] {
      feegrants: {
        num_grantees: 1,
        granter: 'clt16z0herz998946wr659lr84c8c556da55dc34hh', //signer1
        external_granter: false,
        grantees: ['relayer'],
        block_last_verified: 1,
      },
    }] + super.chains[1:],
  },
}

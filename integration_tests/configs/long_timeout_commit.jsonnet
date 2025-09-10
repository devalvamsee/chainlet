local default = import 'default.jsonnet';

default {
  'chainlet_988-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '15s',
      },
    },
    'app-config'+: {
      'blocked-addresses': ['clt16z0herz998946wr659lr84c8c556da55dc34hh'],
    },
  },
}

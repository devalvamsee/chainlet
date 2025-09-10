local config = import 'default.jsonnet';

config {
  'chainlet_988-1'+: {
    validators: [validator {
      client_config: {
        'broadcast-mode': 'sync',
      },
    } for validator in super.validators],
  },
}

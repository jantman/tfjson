#!/usr/bin/env python

import sys
import argparse
import json
import re
from collections import defaultdict

TAG_RE = re.compile(r'^tags\.(\d+)\.(\S+)$')


class TagDiffer:

    def run(self, fpath):
        with open(fpath, 'r') as fh:
            j = json.loads(fh.read())
        for k, v in j['state'].items():
            j['state'][k] = self.fix_asg_tags(dict(v))
        for k, v in j['diff'].items():
            if v == {'destroy': False}:
                continue
            if v['destroy']:
                print(f'- {k}')
                continue
            if v['destroy_tainted']:
                print(f'-/+ {k}')
            else:
                print(f'~ {k}')
            v.pop('destroy')
            v.pop('destroy_tainted')
            v = self.fix_asg_tags(dict(v))
            v.pop('tags.')
            maxlen = max([len(k) for k in v.keys()])
            for attr in sorted(v.keys()):
                print(f'\t{attr:<{maxlen}}: "{j["state"][k][attr]}" => "{v[attr]}"')

    def fix_asg_tags(self, d):
        tags = defaultdict(dict)
        for k in list(d.keys()):
            m = TAG_RE.match(k)
            if not m:
                continue
            tags[m.group(1)][m.group(2)] = d[k]
            del d[k]
        for idx, data in tags.items():
            d[f'tags.{data["key"]}'] = data['value']
        return d


if __name__ == "__main__":
    p = argparse.ArgumentParser(description='tf plan JSON tag diff')
    p.add_argument('PLAN_JSON', type=str, action='store',
                   help='path to tfplan JSON file (tfjson output)')
    args = p.parse_args(sys.argv[1:])
    TagDiffer().run(args.PLAN_JSON)

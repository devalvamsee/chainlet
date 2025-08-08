from pathlib import Path

import pytest

from .network import setup_custom_chainlet
from .utils import submit_any_proposal

pytestmark = pytest.mark.gov


@pytest.fixture(scope="module")
def custom_chainlet(tmp_path_factory):
    path = tmp_path_factory.mktemp("chainlet")
    yield from setup_custom_chainlet(
        path, 26400, Path(__file__).parent / "configs/broadcast.jsonnet"
    )


def test_submit_any_proposal(custom_chainlet):
    submit_any_proposal(custom_chainlet)

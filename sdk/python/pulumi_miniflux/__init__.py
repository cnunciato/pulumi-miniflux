# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .miniflux_service import *
from .provider import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "miniflux",
  "mod": "index",
  "fqn": "pulumi_miniflux",
  "classes": {
   "miniflux:index:MinifluxService": "MinifluxService"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "miniflux",
  "token": "pulumi:providers:miniflux",
  "fqn": "pulumi_miniflux",
  "class": "Provider"
 }
]
"""
)

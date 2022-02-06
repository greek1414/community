"""
Applet: Leetcode
Summary: Programming challenges
Description: Displays the daily Leetcode challenge.
Author: greek1414
"""

load("render.star", "render")
load("schema.star", "schema")

def main(config):
    return render.Root(
        child = render.Box(height = 1, width = 1),
    )

def get_schema():
    return schema.Schema(
        version = "1",
        fields = [],
    )
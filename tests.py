import unittest
from unittest import TestCase
from algs.main import edit_string


class TestEditString(TestCase):
    def test_single_replace(self):
        str1 = 'man'
        str2 = 'num'
        result = edit_string(1, 10.1)[0]
        self.assertEqual(result, 'replace')


if __name__ == '__main__':
    unittest.main()

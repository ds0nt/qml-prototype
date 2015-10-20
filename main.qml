import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0

ApplicationWindow {
  visible: true
  title: "Basic layouts"
  property int margin: 11
  width: mainLayout.implicitWidth + 2 * margin
  height: mainLayout.implicitHeight + 2 * margin
  minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
  minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin

  ColumnLayout {
    id: mainLayout
    anchors.fill: parent
    anchors.margins: margin
    GroupBox {
      id: rowBox
      title: "Row layout"
      Layout.fillWidth: true

      RowLayout {
        id: rowLayout
        anchors.fill: parent
        TextField {
          placeholderText: "Enter Search Query"
          Layout.fillWidth: true
          objectName: "query"
        }
        Button {
          text: "Search"
          objectName: "submit"
        }
      }
    }
    TextArea {
      objectName: "results"
      id: t3
      text: "This fills the whole cell"
      Layout.minimumHeight: 30
      Layout.fillHeight: true
      Layout.fillWidth: true
    }
  }
}
